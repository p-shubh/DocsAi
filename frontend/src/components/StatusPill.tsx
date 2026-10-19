import type { DocumentStatus } from "@/lib/types";

const STYLES: Record<DocumentStatus, string> = {
  uploaded: "bg-hover text-ink-muted",
  queued: "bg-info-soft text-info",
  processing: "bg-info-soft text-info animate-[docsai-pulse_1.4s_ease-in-out_infinite]",
  completed: "bg-success-soft text-success",
  failed: "bg-danger-soft text-danger",
};

export function StatusPill({ status }: { status: DocumentStatus }) {
  return (
    <span
      className={`inline-flex items-center gap-1 rounded-full px-2.5 py-0.5 text-[10.5px] font-semibold capitalize shrink-0 ${STYLES[status]}`}
    >
      {status}
    </span>
  );
}
