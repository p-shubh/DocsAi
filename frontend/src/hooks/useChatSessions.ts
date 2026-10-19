"use client";

import { useCallback, useState } from "react";
import { ApiError } from "@/lib/api";
import { streamAskQuestion } from "@/lib/streamChat";
import type { SourceCitation } from "@/lib/types";

export type ChatMessage =
  | { id: string; kind: "user"; text: string }
  | {
      id: string;
      kind: "assistant";
      text: string;
      sources: SourceCitation[];
    }
  | { id: string; kind: "error"; title: string; text: string; question: string };

export interface ChatSession {
  id: string;
  title: string;
  messages: ChatMessage[];
  loading: boolean;
}

let uid = 0;
function nextId(prefix: string) {
  uid += 1;
  return `${prefix}_${uid}`;
}

function makeSession(): ChatSession {
  return { id: nextId("sess"), title: "New conversation", messages: [], loading: false };
}

export function useChatSessions() {
  const [sessions, setSessions] = useState<ChatSession[]>([makeSession()]);
  const [currentSessionId, setCurrentSessionId] = useState(sessions[0].id);
  const [chatInput, setChatInput] = useState("");
  const [openSources, setOpenSources] = useState<Record<string, boolean>>({});

  const currentSession =
    sessions.find((s) => s.id === currentSessionId) ?? sessions[0];

  const updateSession = useCallback(
    (id: string, updater: (session: ChatSession) => ChatSession) => {
      setSessions((prev) => prev.map((s) => (s.id === id ? updater(s) : s)));
    },
    [],
  );

  const newSession = useCallback(() => {
    const fresh = makeSession();
    setSessions((prev) => [fresh, ...prev]);
    setCurrentSessionId(fresh.id);
    setChatInput("");
  }, []);

  const selectSession = useCallback((id: string) => {
    setCurrentSessionId(id);
    setChatInput("");
  }, []);

  const toggleSources = useCallback((msgId: string) => {
    setOpenSources((prev) => ({ ...prev, [msgId]: !prev[msgId] }));
  }, []);

  const runQuery = useCallback(
    (question: string, sessionId: string) => {
      const assistantId = nextId("m");

      const fail = (title: string, text: string) => {
        updateSession(sessionId, (s) => ({
          ...s,
          messages: [
            ...s.messages.filter((m) => m.id !== assistantId),
            { id: nextId("m"), kind: "error", title, text, question },
          ],
          loading: false,
        }));
      };

      streamAskQuestion(question, undefined, {
        onSources: (sources) => {
          updateSession(sessionId, (s) => ({
            ...s,
            messages: [...s.messages, { id: assistantId, kind: "assistant", text: "", sources }],
          }));
        },
        onDelta: (delta) => {
          updateSession(sessionId, (s) => ({
            ...s,
            loading: false,
            messages: s.messages.map((m) =>
              m.id === assistantId && m.kind === "assistant"
                ? { ...m, text: m.text + delta }
                : m,
            ),
          }));
        },
        onDone: () => {
          updateSession(sessionId, (s) => ({ ...s, loading: false }));
        },
        onError: (message) => fail("Something went wrong", message),
      }).catch((err) => {
        const isUnavailable = err instanceof ApiError && err.status === 502;
        fail(
          isUnavailable ? "AI service unavailable" : "Something went wrong",
          err instanceof ApiError ? err.message : "The request failed unexpectedly.",
        );
      });
    },
    [updateSession],
  );

  const sendMessage = useCallback(() => {
    const question = chatInput.trim();
    if (!question || !currentSession || currentSession.loading) return;

    const sessionId = currentSession.id;
    const userMsg: ChatMessage = { id: nextId("m"), kind: "user", text: question };
    const isFirst = currentSession.messages.length === 0;

    updateSession(sessionId, (s) => ({
      ...s,
      title: isFirst
        ? question.slice(0, 32) + (question.length > 32 ? "…" : "")
        : s.title,
      messages: [...s.messages, userMsg],
      loading: true,
    }));
    setChatInput("");
    runQuery(question, sessionId);
  }, [chatInput, currentSession, runQuery, updateSession]);

  const retry = useCallback(
    (sessionId: string, errorMsgId: string, question: string) => {
      updateSession(sessionId, (s) => ({
        ...s,
        messages: s.messages.filter((m) => m.id !== errorMsgId),
        loading: true,
      }));
      runQuery(question, sessionId);
    },
    [runQuery, updateSession],
  );

  return {
    sessions,
    currentSession,
    chatInput,
    setChatInput,
    openSources,
    newSession,
    selectSession,
    toggleSources,
    sendMessage,
    retry,
  };
}
