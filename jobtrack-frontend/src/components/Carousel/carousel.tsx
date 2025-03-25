import React, { useRef, useState } from "react";
import ApplicationCard from "../Cards/applicationCard";
import { Application, Status } from "../../types/application";

import "./carousel.css"

interface CarouselProps {
    items: Application[];
    onUpdate: (id: number, newStatus: Status) => void;
    onDelete: (id: number) => void;
}

const Carousel: React.FC<CarouselProps> = ({ items, onUpdate, onDelete }) => {
    const carouselRef = useRef<HTMLDivElement>(null);
    const [selectedIndex, setSelectedIndex] = useState<number>(0);

    const handleSelect = (index: number) => {
        setSelectedIndex(index);
    };

    return (
        <div className="carousel-wrapper">
            <div className="carousel-container" ref={carouselRef}>
                {items.map((application, index) => (
                    <div key={index} className="card-wrapper">
                        <ApplicationCard 
                            application={application}
                            className={selectedIndex === index ? "selected" : ""}
                            isSelected={index === selectedIndex}
                            onClick={() => handleSelect(index)}
                            onUpdate={onUpdate}
                            onDelete={onDelete}
                        />
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Carousel;