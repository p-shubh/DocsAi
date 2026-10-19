"use client";

import type { ChatSession } from "@/hooks/useChatSessions";

export function SessionsRail({
  sessions,
  currentSessionId,
  onNewSession,
  onSelectSession,
}: {
  sessions: ChatSession[];
  currentSessionId: string;
  onNewSession: () => void;
  onSelectSession: (id: string) => void;
}) {
  return (
    <div className="flex w-[220px] shrink-0 flex-col gap-2.5 border-r border-border p-3">
      <button
        type="button"
        onClick={onNewSession}
        className="rounded-lg border border-border bg-surface px-3 py-2 text-center text-[13px] font-semibold text-ink"
      >
        + New chat
      </button>
      <div className="flex flex-col gap-0.5 overflow-y-auto">
        {sessions.map((sess) => {
          const active = sess.id === currentSessionId;
          const qCount = sess.messages.filter((m) => m.kind === "user").length;
          return (
            <button
              type="button"
              key={sess.id}
              onClick={() => onSelectSession(sess.id)}
              className={`rounded-md px-2.5 py-2 text-left ${
                active ? "bg-accent-soft" : "hover:bg-hover"
              }`}
            >
              <div
                className={`truncate text-[12.5px] font-semibold ${
                  active ? "text-accent" : "text-ink"
                }`}
              >
                {sess.title}
              </div>
              <div className="mt-0.5 text-[11px] text-ink-faint">
                {qCount ? `${qCount} question${qCount === 1 ? "" : "s"}` : "Empty"}
              </div>
            </button>
          );
        })}
      </div>
    </div>
  );
}
