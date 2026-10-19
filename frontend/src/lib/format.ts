export function fmtSize(bytes: number): string {
  if (bytes < 1024) return `${bytes} B`;
  const units = ["KB", "MB", "GB"];
  let val = bytes / 1024;
  let i = 0;
  while (val >= 1024 && i < units.length - 1) {
    val /= 1024;
    i++;
  }
  return `${val.toFixed(val < 10 ? 1 : 0)} ${units[i]}`;
}

export function fmtRelTime(iso: string): string {
  const diffMs = Date.now() - new Date(iso).getTime();
  const mins = Math.floor(diffMs / 60000);
  if (mins < 1) return "just now";
  if (mins < 60) return `${mins}m ago`;
  const hrs = Math.floor(mins / 60);
  if (hrs < 24) return `${hrs}h ago`;
  const days = Math.floor(hrs / 24);
  return `${days}d ago`;
}

const MIME_EXT_MAP: Record<string, string> = {
  "application/pdf": "PDF",
  "text/plain": "TXT",
  "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
    "DOCX",
  "text/csv": "CSV",
  "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
    "XLSX",
  "text/markdown": "MD",
};

export function extFromMime(mime: string, name: string): string {
  if (MIME_EXT_MAP[mime]) return MIME_EXT_MAP[mime];
  const parts = name.split(".");
  return parts.length > 1
    ? (parts.pop() ?? "").slice(0, 4).toUpperCase()
    : "FILE";
}
