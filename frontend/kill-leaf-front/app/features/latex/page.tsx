import { LatexEditor } from "./components/LatexEditor";

export default function LatexPage() {
  return (
    <main className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">LaTeX Editor</h1>
      <LatexEditor />
    </main>
  );
}
