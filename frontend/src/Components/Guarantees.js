import React from "react";

export default class Guarantees extends React.Component {
    render() {
        return (
            <div className='guarantees_wrapper'>
  
                <div className='guarantee'>
                  <div className='guarantee_img'>
                    <picture>
                      <img src='./images/guarantees/Гарантия1.webp' alt='Guarantee1' />
                    </picture>
                  </div>
                
                  <div className='guarantee_text'>
                    Гарантия на <br/> изготавливаемые <br/> изделия — <span>2 года</span>
                  </div>
                </div>
            
                <div className='guarantee'>
                  <div className='guarantee_img'>
                    <picture>
                      <img src='./images/guarantees/Размеры1.webp' alt='Guarantee2' />
                    </picture>
                  </div>
            
                  <div className='guarantee_text'>
                    Любую модель изготовим <br/> <span>на заказ</span> по вашим <br/> размерам
                  </div>
                </div>
                
                <div className='guarantee'>
                  <div className='guarantee_img'>
                    <picture>
                      <img src='./images/guarantees/Рассрочка1.webp' alt='Guarantee3' />
                    </picture>
                  </div>
            
                  <div className='guarantee_text'>
                    <span>Беспроцентная</span> <br/> рассрочка <br/> на 12 месяцев
                  </div>
                </div>
            
                <div className='guarantee'>
                  <div className='guarantee_img'>
                    <picture>
                      <img src='./images/guarantees/Сборка1.webp' alt='Guarantee4' />
                    </picture>
                  </div>
            
                  <div className='guarantee_text'>
                    <span>Бесплатный подъём</span> <br/> на этаж, уборка <br/> упаковочных матриалов
                  </div>
                </div>
  
            </div>
        );
    }
}