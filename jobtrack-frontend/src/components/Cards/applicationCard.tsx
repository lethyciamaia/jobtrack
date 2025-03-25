import React, { useEffect, useState } from "react";
import { format } from 'date-fns';
import { VStack, HStack, Spacer } from "../Stacks/stack";
import { Application, Status } from "../../types/application";

import "./card.css"
import "./dropdown.css"

interface StatusTagProps {
    status: Status;  
}

const StatusTag: React.FC<StatusTagProps> = ({ status }) => {
    const statusClass = `status-tag status-${status.toLowerCase()}`;

    return <span className={statusClass}>{status}</span>;
}

type ApplicationCardProps = {
    application: Application;
    className?: string; 
    isSelected: boolean;
    onClick: () => void;
    onUpdate: (id: number, newStatus: Status) => void;
    onDelete: (id: number) => void;
};

const ApplicationCard: React.FC<ApplicationCardProps> = ({ application, className = "", isSelected, onClick, onUpdate, onDelete }) => {
    const [showMenu, setShowMenu] = useState(false);
    console.log(application);

    const handleCardClick = () => {
        onClick(); 
        setShowMenu(false); 
    };

    const handleStatusClick = (e: React.MouseEvent) => {
        e.stopPropagation();
        if (isSelected) {
            setShowMenu(!showMenu);  
        } else {
            onClick(); 
        }
    };

    const handleDelete = (e: React.MouseEvent) => {
        e.stopPropagation(); 
        if (isSelected) {
            onDelete(application.id); 
        } else {
            onClick(); 
        }
    };

    // useEffect()

    return (
        <div 
            onClick={handleCardClick}
            style={ {
                cursor:'pointer',
                transition: 'transform 0.3s ease, opacity 0.3s ease',
                transform: isSelected ? 'scale(1.2)' : 'scale(1)',
                opacity: isSelected ? 1 : 0.7,
                margin: '0 15px',
                boxShadow: isSelected ? '0 8px 16px rgba(0, 0, 0, 0.2)' : 'none',
            }}
            className={`card-container ${className}`}
        > 
            <VStack > 
                <HStack gap={4}> 
                    <VStack gap={10}>
                        <h1>{application.company}</h1>
                        <h3>Created at {application.created_at ? format(new Date(application.created_at), 'dd/MM/yyyy') : 'Invalid Date'}</h3>
                    </VStack>
                    <Spacer />
                    <div onClick={handleStatusClick}>
                        <StatusTag status={application.status}/>
                    </div>
                    {showMenu && (
                        <ul className="dropdown-menu">
                        {Object.values(Status).map((status) => (
                        <li key={status} onClick={() => {
                            onUpdate(application.id, status); 
                            setShowMenu(false);
                        }}>
                            {status}
                        </li>
                        ))}
                        </ul>
                    )}
                </HStack>
                <p> {application.description} </p>
                <Spacer />
                <div onClick={handleDelete} className="delete-button">
                    <h5> DELETE </h5>
                </div>
            </VStack>
        </div>
    );
};

export default ApplicationCard;