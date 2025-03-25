import { useState } from "react";
import { HStack, Spacer } from "../Stacks/stack";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus } from '@fortawesome/free-solid-svg-icons';
import ApplicationForm from "../Forms/applicationForm";
import "./header.css"

interface HeaderProps {
    addApplication: (newApplication: any) => void;
}

const Header: React.FC<HeaderProps> = ({ addApplication }) => {
    const [showForm, setShowForm] = useState(false);

    return (
        <div className="header-container"> 
            <HStack >
                <h1>JobTrack</h1>
                <Spacer />
                <div className="add-button" onClick={() => setShowForm(true)}>
                    <FontAwesomeIcon icon={faPlus} size="2x" />
                    <h3> Add application </h3>
                </div>
            </HStack>
            {showForm && (
                <ApplicationForm 
                    onClose={() => setShowForm(false)} 
                    onSubmit={addApplication} 
                />
            )}
        </div>
    )
};

export default Header;