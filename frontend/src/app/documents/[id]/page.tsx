import { DocumentDetailView } from "@/components/documents/DocumentDetailView";

export default async function DocumentDetailPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;
  return <DocumentDetailView documentId={id} />;
}
