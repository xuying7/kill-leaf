"use client";

import { useState } from "react";
import Latex from "react-latex-next";
import "katex/dist/katex.min.css";
import { LatexPreview } from "../components/LatexPreview";
import { LatexToolbar } from "../components/LatexToolbar";

export const LatexEditor = () => {
  const [input, setInput] = useState("");

  return (
    <div className="flex flex-col gap-4">
      <LatexToolbar />
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div className="border rounded-lg p-4">
          <textarea
            className="w-full min-h-[300px] p-2 border rounded"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Enter LaTeX here..."
          />
        </div>
        <LatexPreview content={input} />
      </div>
    </div>
  );
};
