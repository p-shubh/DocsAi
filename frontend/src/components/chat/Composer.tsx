"use client";

export function Composer({
  value,
  disabled,
  sendDisabled,
  onChange,
  onSend,
}: {
  value: string;
  disabled: boolean;
  sendDisabled: boolean;
  onChange: (value: string) => void;
  onSend: () => void;
}) {
  return (
    <div className="flex gap-2.5 border-t border-border bg-bg px-8 py-4">
      <input
        placeholder="Ask a question about your documents..."
        value={value}
        disabled={disabled}
        onChange={(e) => onChange(e.target.value)}
        onKeyDown={(e) => {
          if (e.key === "Enter" && !sendDisabled) onSend();
        }}
        className="flex-1 rounded-[9px] border border-border bg-surface px-3.5 py-2.5 text-sm text-ink outline-none disabled:opacity-60"
      />
      <button
        type="button"
        onClick={onSend}
        disabled={sendDisabled}
        className={`rounded-[9px] px-5 text-[13.5px] font-semibold ${
          sendDisabled
            ? "bg-hover text-ink-faint"
            : "bg-accent text-white"
        }`}
      >
        Send
      </button>
    </div>
  );
}
