import { API_BASE, buildApiError } from "./api";
import type { SourceCitation } from "./types";

export interface StreamCallbacks {
  onSources: (sources: SourceCitation[]) => void;
  onDelta: (text: string) => void;
  onDone: (model: string) => void;
  onError: (message: string) => void;
}

type SSEPayload =
  | { type: "sources"; sources?: SourceCitation[] }
  | { type: "delta"; text: string }
  | { type: "done"; model: string }
  | { type: "error"; message: string };

export async function streamAskQuestion(
  question: string,
  documentId: string | undefined,
  callbacks: StreamCallbacks,
): Promise<void> {
  const res = await fetch(`${API_BASE}/api/v1/chat/stream`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(
      documentId ? { question, document_id: documentId } : { question },
    ),
  });

  if (!res.ok || !res.body) {
    throw await buildApiError(res);
  }

  const reader = res.body.getReader();
  const decoder = new TextDecoder();
  let buffer = "";

  while (true) {
    const { done, value } = await reader.read();
    if (done) break;

    buffer += decoder.decode(value, { stream: true });

    let sepIndex: number;
    while ((sepIndex = buffer.indexOf("\n\n")) !== -1) {
      const rawEvent = buffer.slice(0, sepIndex);
      buffer = buffer.slice(sepIndex + 2);

      const dataLine = rawEvent.split("\n").find((line) => line.startsWith("data:"));
      if (!dataLine) continue;

      const payload = JSON.parse(dataLine.slice(5).trim()) as SSEPayload;

      switch (payload.type) {
        case "sources":
          callbacks.onSources(payload.sources ?? []);
          break;
        case "delta":
          callbacks.onDelta(payload.text);
          break;
        case "done":
          callbacks.onDone(payload.model);
          break;
        case "error":
          callbacks.onError(payload.message);
          break;
      }
    }
  }
}
