import React from "react";
import './stack.css'

type StackProps = {
    children: React.ReactNode;
    gap?: number;
    className?: string;
};

export function VStack({ children, gap = 4, className = "" }: StackProps) {
    return (
    <div 
        className={`vstack ${className}`} 
        style={{ gap: `${gap}px` }}
    >
        { children }
    </div>)
}

export function HStack({ children, gap = 4, className = "" }: StackProps) {
    return (
    <div 
        className={`hstack ${className}`} 
        style={{ gap: `${gap}px` }}
    >
        { children }
    </div>)
}

// Works like Spacer from SwiftUI
export function Spacer() {
    return <div className="spacer" />;
}