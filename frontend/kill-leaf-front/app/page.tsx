"use client";

import "katex/dist/katex.min.css";
import Latex from "react-latex-next";
import { useState } from "react";

export default function Home() {
  const [input, setInput] = useState("");

  return (
    <main className="min-h-screen p-8">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">LaTeX Editor</h1>

        {/* Editor and Preview Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {/* Input Section */}
          <div className="border rounded-lg p-4 bg-white shadow-sm">
            <h2 className="text-lg font-semibold mb-2">Input</h2>
            <textarea
              className="w-full min-h-[300px] p-3 border rounded-md"
              value={input}
              onChange={(e) => setInput(e.target.value)}
              placeholder="Enter your LaTeX here... (e.g., $x^2 + y^2 = r^2$)"
            />
          </div>

          {/* Preview Section */}
          <div className="border rounded-lg p-4 bg-white shadow-sm">
            <h2 className="text-lg font-semibold mb-2">Preview</h2>
            <div className="latex-preview p-3 bg-gray-50 rounded-md min-h-[300px]">
              <Latex>{input}</Latex>
            </div>
          </div>
        </div>

        {/* Quick Reference */}
        <div className="mt-8 p-4 border rounded-lg bg-white">
          <h2 className="text-lg font-semibold mb-2">Quick Reference</h2>
          <p className="text-sm text-gray-600">
            Inline math: $x^2$ <br />
            Display math: $$x^2$$ <br />
            Greek letters: $\alpha, \beta, \gamma$ <br />
            Fractions: $\frac{1}
            {2}$ <br />
            Subscripts: $x_n$ <br />
            Superscripts: $x^n$
          </p>
        </div>
      </div>
    </main>
  );
}
