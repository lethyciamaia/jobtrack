import React from "react";
import { format } from 'date-fns';
import { VStack, HStack, Spacer } from "../Stacks/stack";
import { Application, Status } from "../../types/application";

import "./card.css"

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
};

const ApplicationCard: React.FC<ApplicationCardProps> = ({ application, className = "", isSelected, onClick }) => {
    return (
        <div 
            onClick={onClick}
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
            <VStack className="child"> 
                <HStack gap={4} className=""> 
                    <VStack gap={0}>
                        <h1>{application.enterprise}</h1>
                        <h3>Created at {application.createdAt ? format(new Date(application.createdAt), 'dd/MM/yyyy') : 'Invalid Date'}</h3>

                    </VStack>
                    <Spacer />
                    <StatusTag status={application.status}/>
                </HStack>
                <p> {application.description} </p>
            </VStack>
        </div>
    );
};

export default ApplicationCard;