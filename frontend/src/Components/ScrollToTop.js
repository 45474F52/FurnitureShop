import React, { useEffect, useRef } from 'react'

export default function ScrollToTop() {
    const element = useRef(null);

    useEffect(() => {
        window.scrollTo(0, 0);
    });

    return (
        <div ref={element} style={{display: 'none' }}></div>
    );
}