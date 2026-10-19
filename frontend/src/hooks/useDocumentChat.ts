"use client";

import { useCallback, useState } from "react";
import { ApiError } from "@/lib/api";
import { streamAskQuestion } from "@/lib/streamChat";
import type { ChatMessage } from "./useChatSessions";

let uid = 0;
function nextId(prefix: string) {
  uid += 1;
  return `${prefix}_${uid}`;
}

export function useDocumentChat(documentId: string) {
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const [loading, setLoading] = useState(false);
  const [input, setInput] = useState("");
  const [openSources, setOpenSources] = useState<Record<string, boolean>>({});

  const toggleSources = useCallback((msgId: string) => {
    setOpenSources((prev) => ({ ...prev, [msgId]: !prev[msgId] }));
  }, []);

  const runQuery = useCallback(
    (question: string) => {
      const assistantId = nextId("m");

      const fail = (title: string, text: string) => {
        setMessages((prev) => [
          ...prev.filter((m) => m.id !== assistantId),
          { id: nextId("m"), kind: "error", title, text, question },
        ]);
        setLoading(false);
      };

      streamAskQuestion(question, documentId, {
        onSources: (sources) => {
          setMessages((prev) => [
            ...prev,
            { id: assistantId, kind: "assistant", text: "", sources },
          ]);
        },
        onDelta: (delta) => {
          setLoading(false);
          setMessages((prev) =>
            prev.map((m) =>
              m.id === assistantId && m.kind === "assistant"
                ? { ...m, text: m.text + delta }
                : m,
            ),
          );
        },
        onDone: () => setLoading(false),
        onError: (message) => fail("Something went wrong", message),
      }).catch((err) => {
        const isUnavailable = err instanceof ApiError && err.status === 502;
        fail(
          isUnavailable ? "AI service unavailable" : "Something went wrong",
          err instanceof ApiError ? err.message : "The request failed unexpectedly.",
        );
      });
    },
    [documentId],
  );

  const sendMessage = useCallback(() => {
    const question = input.trim();
    if (!question || loading) return;

    const userMsg: ChatMessage = { id: nextId("m"), kind: "user", text: question };
    setMessages((prev) => [...prev, userMsg]);
    setInput("");
    setLoading(true);
    runQuery(question);
  }, [input, loading, runQuery]);

  const retry = useCallback(
    (errorMsgId: string, question: string) => {
      setMessages((prev) => prev.filter((m) => m.id !== errorMsgId));
      setLoading(true);
      runQuery(question);
    },
    [runQuery],
  );

  return {
    messages,
    loading,
    input,
    setInput,
    openSources,
    toggleSources,
    sendMessage,
    retry,
  };
}
