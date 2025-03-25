import React, { useState } from "react";
import "./applicationForm.css"

interface ApplicationFormProps {
    onClose: () => void;
    onSubmit: (newApplication: any) => void;
}

const ApplicationForm: React.FC<ApplicationFormProps> = ({ onClose, onSubmit }) => {
    const [formData, setFormData] = useState({
        company: "",
        createdAt: new Date().toISOString(),
        status: "Sended",
        description: "",
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        onSubmit({ id: Date.now(), ...formData });
        onClose();
    };

    return (
        <div className="form-overlay">
            <div className="form-container">
                <h2>New application</h2>
                <form onSubmit={handleSubmit}>
                    <label>Company:</label>
                    <input type="text" name="company" value={formData.company} onChange={handleChange} required />

                    <label>Description:</label>
                    <textarea name="description" value={formData.description} onChange={handleChange} required />

                    <button type="submit">Add</button>
                    <button type="button" onClick={onClose}>Cancel</button>
                </form>
            </div>
        </div>
    );
};

export default ApplicationForm;
