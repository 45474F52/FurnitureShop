import React, { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import ScrollToTop from '../Components/ScrollToTop';

export default function Furnitures() {
    const serverUrl = process.env.REACT_APP_SERVER_URL;

    let { tag } = useParams();
    const [data, setData] = useState([]);

    useEffect(() => {
        fetchData();
    }, []);

    const fetchData = async () => {
        try {
            const response = await fetch(serverUrl + '/furnitures?category=' + tag);
            const json = await response.json();
            setData(json);
        } catch (error) {
            console.error('Ошибка при получении товаров: ', error);
        }
    };

    return (
        <div>
            <ScrollToTop />

            <div className='furnitures-header'>
                <a href='/' />
                <h3>{tag}</h3>
            </div>

            <div className='furnitures-container'>
                {
                    data.map((item) => (
                        <div key={item.id} className='furniture'>
                            <img src={serverUrl + '/images?image=' + item.image.String} alt={item.name} />
                            <h3>{item.name}</h3>
                            <p>{item.price} ₽</p>
                        </div>
                    ))
                }
            </div>
        </div>
    );
}