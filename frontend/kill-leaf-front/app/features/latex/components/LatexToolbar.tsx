"use client";

export const LatexToolbar = () => {
  return (
    <div className="flex gap-2 p-2 border rounded-lg bg-gray-50">
      <button className="px-3 py-1 rounded hover:bg-gray-200">$...$</button>
      <button className="px-3 py-1 rounded hover:bg-gray-200">$$...$$</button>
      <button className="px-3 py-1 rounded hover:bg-gray-200">\begin{}</button>
      {/* Add more LaTeX shortcuts */}
    </div>
  );
};
