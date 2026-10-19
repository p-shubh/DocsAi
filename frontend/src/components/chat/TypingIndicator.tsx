export function TypingIndicator() {
  return (
    <div className="mx-auto mb-5 max-w-[720px] rounded-xl border border-border bg-surface px-4 py-3.5">
      <div className="mb-1.5 text-[11px] font-bold uppercase tracking-wide text-accent">
        DocsAI
      </div>
      <div className="flex gap-1 py-0.5">
        {[0, 150, 300].map((delay) => (
          <span
            key={delay}
            style={{ animationDelay: `${delay}ms` }}
            className="h-1.5 w-1.5 animate-[docsai-blink_1s_ease-in-out_infinite] rounded-full bg-ink-faint"
          />
        ))}
      </div>
    </div>
  );
}
