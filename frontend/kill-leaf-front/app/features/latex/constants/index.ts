export const LATEX_DELIMITERS = [
    { left: '$$', right: '$$', display: true },
    { left: '\\(', right: '\\)', display: false },
    { left: '$', right: '$', display: false },
    { left: '\\[', right: '\\]', display: true },
  ];
  
  export const LATEX_MACROS = {
    "\\f": "#1f(#2)",
  };
  
  export const COMMON_LATEX_SNIPPETS = [
    { label: 'Inline Math', snippet: '$x^2$' },
    { label: 'Display Math', snippet: '$$\\frac{-b \\pm \\sqrt{b^2-4ac}}{2a}$$' },
    { label: 'Equation', snippet: '\\begin{equation}\n\\end{equation}' },
  ];