import React, { useState, useEffect } from 'react'
import { useNavigate } from "react-router-dom";
import Guarantees from '../Components/Guarantees';

export default function StartPage() {
    const navigate = useNavigate();
    const serverUrl = process.env.REACT_APP_SERVER_URL;
    const [data, setData] = useState([]);

    useEffect(() => {
        fetchData();
    }, []);

    const fetchData = async () => {
        try {
            const response = await fetch(serverUrl + '/categories');
            const json = await response.json();
            setData(json);
        } catch (error) {
            console.error('Ошибка при получении категорий товаров: ', error);
        }
    };

    return (
        <div>
            <div className='nav_cards'>

                {
                    data.map((item) => (
                        <div
                            key={item.id}
                            className='nav_card'
                            onClick={() => { navigate('/furnitures/' + item.Name); }}>
                                <img src={serverUrl + '/images?image=' + item.image.String} alt={item.Name} />
                                <h2>{item.Name}</h2>
                        </div>
                    ))
                }
                
            </div>

            <Guarantees />
        </div>
    );
}