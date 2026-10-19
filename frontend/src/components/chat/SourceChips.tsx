import type { SourceCitation } from "@/lib/types";

export function SourceChips({ sources }: { sources: SourceCitation[] }) {
  return (
    <div className="mt-2 flex flex-wrap gap-1.5">
      {sources.map((src) => {
        const location = [src.page ? `p.${src.page}` : null, src.section]
          .filter(Boolean)
          .join(" · ");
        return (
          <div
            key={src.chunk_id}
            className="rounded-full border border-border bg-hover px-2.5 py-1 text-[11.5px] text-ink-muted"
          >
            <span className="font-semibold">{src.document_name}</span>
            {location && <span className="opacity-65"> · {location}</span>}
          </div>
        );
      })}
    </div>
  );
}
