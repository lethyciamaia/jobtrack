// import { useState, useEffect } from "react";
import { HStack, Spacer } from "../Stacks/stack";
import "./header.css"

const Header = () => {
    // const [scrolled, setScrolled] = useState(false);

    // useEffect(() => {
    //     const handleScroll = () => {
    //         if (window.scrollY > 50) {
    //             setScrolled(true);
    //         } else {
    //             setScrolled(false);
    //         }
    //     };

    //     window.addEventListener("scroll", handleScroll);
    //     return () => window.removeEventListener("scroll", handleScroll);
    // }, []);

    return (
        <div className="header-container"> 
            <HStack >
                <h1>JobTrack</h1>
                <Spacer />
                <h3> filtro de busca </h3>
            </HStack>
        </div>
    )
};

export default Header;