export interface LatexPreviewProps {
    content: string;
  }
  
  export interface LatexDocument {
    id?: string;
    content: string;
    title?: string;
    createdAt?: Date;
    updatedAt?: Date;
  }