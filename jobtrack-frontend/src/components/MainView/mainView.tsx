import React, { useState, useEffect} from "react";
import { Application } from "../../types/application";
import Carousel from "../Carousel/carousel";
import "./mainView.css"

function MainView() {
    const [applications, setApplications] = useState<Application[]>([]);

    useEffect(() => {
        const fetchApplications = async () => {
            try {
                // const apiUrl = process.env.REACT_APP_API_URL;
                // // console.log('URL da api: ',apiUrl);
                // const response = await fetch(`${apiUrl}/applications`);
                const response = await fetch("http://localhost:8000/applications");
                console.log('Resposta da api: ',response);
                const data = await response.json();
                setApplications(data);
            } catch (error) {
                console.error('Failed to fetch applications', error);
            }
        };
        fetchApplications();
    }, []);

    return (
        <div className="main-content"> 
            <Carousel items={applications} />
        </div>
    );
};

export default MainView;