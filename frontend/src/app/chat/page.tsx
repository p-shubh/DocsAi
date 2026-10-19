"use client";

import { Composer } from "@/components/chat/Composer";
import { MessageBubble } from "@/components/chat/MessageBubble";
import { SessionsRail } from "@/components/chat/SessionsRail";
import { TypingIndicator } from "@/components/chat/TypingIndicator";
import { useChatSessions } from "@/hooks/useChatSessions";

const EXAMPLE_QUESTIONS = [
  "What are the key takeaways from these documents?",
  "Summarize the main points.",
  "Is there anything about pricing?",
];

export default function ChatPage() {
  const {
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
  } = useChatSessions();

  const isEmpty = currentSession.messages.length === 0 && !currentSession.loading;
  const sendDisabled = currentSession.loading || !chatInput.trim();

  return (
    <div className="flex flex-1 min-h-0">
      <SessionsRail
        sessions={sessions}
        currentSessionId={currentSession.id}
        onNewSession={newSession}
        onSelectSession={selectSession}
      />

      <div className="flex flex-1 min-w-0 flex-col">
        <div className="border-b border-border px-8 pb-3.5 pt-5">
          <div className="text-[22px] font-bold tracking-tight text-ink">Chat</div>
          <div className="mt-1 text-[13.5px] text-ink-muted">
            Ask anything about your uploaded documents — DocsAI searches your whole
            library.
          </div>
        </div>

        <div className="flex-1 overflow-y-auto px-8 py-6">
          {isEmpty && (
            <div className="mx-auto mt-14 max-w-[560px] text-center">
              <div className="mx-auto flex h-10 w-10 items-center justify-center rounded-[10px] bg-hover text-lg text-ink-faint">
                ?
              </div>
              <div className="mt-3.5 text-[14.5px] font-semibold text-ink">
                Ask your first question
              </div>
              <div className="mt-1 flex flex-wrap justify-center gap-2">
                {EXAMPLE_QUESTIONS.map((q) => (
                  <button
                    type="button"
                    key={q}
                    onClick={() => setChatInput(q)}
                    className="rounded-full border border-border bg-surface px-3 py-1.5 text-[12.5px] text-ink-muted"
                  >
                    {q}
                  </button>
                ))}
              </div>
            </div>
          )}

          {currentSession.messages.map((msg) => (
            <MessageBubble
              key={msg.id}
              message={msg}
              sourcesOpen={!!openSources[msg.id]}
              onToggleSources={() => toggleSources(msg.id)}
              onRetry={() =>
                msg.kind === "error" &&
                retry(currentSession.id, msg.id, msg.question)
              }
            />
          ))}

          {currentSession.loading && <TypingIndicator />}
        </div>

        <Composer
          value={chatInput}
          disabled={currentSession.loading}
          sendDisabled={sendDisabled}
          onChange={setChatInput}
          onSend={sendMessage}
        />
      </div>
    </div>
  );
}
