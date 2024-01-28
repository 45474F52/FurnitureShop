import React from 'react'
import { useNavigate } from "react-router-dom";
import Guarantees from '../Components/Guarantees';

export default function StartPage() {
    const navigate = useNavigate();

    return (
        <div>
            <div className='nav_cards'>

                <div className='nav_card' onClick={() => { navigate('/furnitures/sofa'); }}>
                    <img src='./images/Sofa.jpg' alt="Диваны" />
                    <h2>Диваны</h2>
                </div>

                <div className='nav_card' onClick={() => { navigate('/furnitures/bed'); }}>
                    <img src='./images/Bed.jpg' alt='Кровати' />
                    <h2>Кровати</h2>
                </div>
            </div>

            <Guarantees />
        </div>
    );
}