"use client";

import Latex from "react-latex-next";
import { LatexPreviewProps } from "../types/index";

export const LatexPreview = ({ content }: LatexPreviewProps) => {
  return (
    <div className="border rounded-lg p-4 bg-white">
      <h3 className="text-lg font-semibold mb-2">Preview</h3>
      <div className="latex-preview">
        <Latex>{content}</Latex>
      </div>
    </div>
  );
};
