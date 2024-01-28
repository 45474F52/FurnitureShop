import React from 'react'
import { useParams } from 'react-router-dom'
import ScrollToTop from '../Components/ScrollToTop';

export default function Furnitures() {
    let { tag } = useParams()
    const path = tag === 'sofa' ? '../../furniture_images/sofas' : '../../furniture_images/beds'

    return (
        <div>
            <ScrollToTop />

            <div style={{ marginTop: 150 + 'px' }}>
                <h1>{tag}</h1>
            </div>
        </div>
    );
}