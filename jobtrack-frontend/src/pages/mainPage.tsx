import React, { useState, useEffect} from "react";
import { Application, Status } from "../types/application";
import { fetchApplications, createApplication, updateStatus, deleteApplication } from "../service/applicationService"
import Header from "../components/Header/header";
import Carousel from "../components/Carousel/carousel";
import { VStack } from "../components/Stacks/stack";

import "./mainPage.css"

const MainPage:React.FC = () => {
    const [applications, setApplications] = useState<Application[]>([]);

    useEffect(() => {
        const loadApplications = async () => {
            try {
                const data = await fetchApplications();
                setApplications(data);
            } catch (error) {
                console.error('Failed to load applications', error);
            }
        };
        loadApplications();
    }, []);

    const handleCreate = async (newApplication: Application) => {
        try {
            const createdApplication = await createApplication(newApplication);
            setApplications((prev) => [
                ...prev,
                createdApplication,
            ]);
        } catch (error) {
            
        }
    }
    const handleUpdate = async (id: number, newStatus: Status) => {
        try {
            const updatedApp = await updateStatus(id, newStatus);
            setApplications((prev) =>
                prev.map((app) =>
                    app.id === id ? updatedApp : app
                )
            );
        } catch (error) {
            console.error("Error deleting application", error);
        }
    }
    const handleDelete = async (id: number) => {
        try {
            await deleteApplication(id);
            setApplications((prev) =>
                prev.filter((app) => app.id !== id)
            );
        } catch (error) {
            console.error("Error deleting application", error);
        }
    }

    return (
        <VStack>
            <Header addApplication={ handleCreate }/>
            <div className="carousel"> 
                <Carousel 
                    items={applications} 
                    onUpdate={handleUpdate}
                    onDelete={handleDelete} 
                />
            </div>
        </VStack>
    );
};

export default MainPage;