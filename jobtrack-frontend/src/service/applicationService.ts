import { Application, Status } from "../types/application"

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8000";

export const fetchApplications = async (): Promise<Application[]> => {
    try {
        const response = await fetch(`${API_URL}/applications`);
        if (!response.ok) {
            throw new Error("Failed to fetch applications")
        }
        return await response.json();
    } catch (error) {
        console.error('Error fetching applications:', error);
        throw error;
    }
};

export const createApplication = async (newApplication:Application): Promise<Application> => {
    try {
        const response = await fetch(`${API_URL}/applications`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(newApplication),
        });

        if (!response.ok) {
            throw new Error("Failed to create application");
        }
        return await response.json();
    } catch (error) {
        console.error('Error creating application:', error);
        throw error;
    }
}

export const updateStatus = async (id:number, newStatus: Status): Promise<Application> => {
    try { 
        const response = await fetch(`${API_URL}/applications`, {
            method: 'PATCH',
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({id: id, status: newStatus}),
        });

        if (!response.ok) {
            throw new Error("Failed to update status")
        }
        return await response.json();
    } catch (error) {
        console.error("Error updating status:", error)
        throw error;
    }
}

export const deleteApplication = async (id: number): Promise<void> => {
    try {
        const response = await fetch(`${API_URL}/applications/${id}`, {
            method: 'DELETE',
        });
        if (!response.ok) {
            throw new Error("Failed to delete application")
        }
    } catch (error) {
        console.error("Error deleting application:", error)
        throw error;
    }
}