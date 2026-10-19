import type { ChatMessage } from "@/hooks/useChatSessions";
import { SourceChips } from "./SourceChips";

export function MessageBubble({
  message,
  sourcesOpen,
  onToggleSources,
  onRetry,
}: {
  message: ChatMessage;
  sourcesOpen: boolean;
  onToggleSources: () => void;
  onRetry: () => void;
}) {
  if (message.kind === "user") {
    return (
      <div className="mb-4 flex justify-end">
        <div className="max-w-[70%] rounded-[14px] rounded-br-[3px] bg-accent px-3.5 py-2.5 text-sm leading-relaxed text-white">
          {message.text}
        </div>
      </div>
    );
  }

  if (message.kind === "error") {
    return (
      <div className="mx-auto mb-5 max-w-[720px] rounded-xl border border-danger-soft bg-danger-soft px-4 py-3.5">
        <div className="text-[13.5px] font-semibold text-danger">
          {message.title}
        </div>
        <div className="mt-0.5 text-[13px] text-ink-muted">{message.text}</div>
        <button
          type="button"
          onClick={onRetry}
          className="mt-2.5 rounded-md border border-danger px-3 py-1.5 text-[12.5px] font-semibold text-danger"
        >
          Retry
        </button>
      </div>
    );
  }

  return (
    <div className="mx-auto mb-5 max-w-[720px] rounded-xl border border-border bg-surface px-4 py-3.5">
      <div className="mb-1.5 text-[11px] font-bold uppercase tracking-wide text-accent">
        DocsAI
      </div>
      <div className="whitespace-pre-wrap text-[14.5px] leading-relaxed text-ink">
        {message.text}
      </div>
      {message.sources.length > 0 && (
        <>
          <button
            type="button"
            onClick={onToggleSources}
            className="mt-2.5 inline-block text-[11.5px] font-semibold text-ink-muted"
          >
            {sourcesOpen
              ? `▾ Hide sources (${message.sources.length})`
              : `▸ Show sources (${message.sources.length})`}
          </button>
          {sourcesOpen && <SourceChips sources={message.sources} />}
        </>
      )}
    </div>
  );
}
