import React from 'react';

export default class Footer extends React.Component {
    render() {
        return (
            <footer>
                <div className='about_us'>
                    <div className='logo_block'>
                        Магазин Мебели
                    </div>

                    <div className='info_block'>
                        <div className='tel_wrapper'>
                            <div className='tel'>
                                <a href='tel:88005552316'>8 800 555 23 16</a>
                            </div>
                            <div className='tel'>
                                <a href='tel:84955405086'>8 495 540 50 86</a>
                            </div>
                        </div>

                        <div className='time_wrapper'>
                            <div>Будни с 10 до 21</div>
                            <div>Выходные с 11 до 19</div>
                        </div>

                        <div className='email_wrapper'>
                            <a href='mailto:sale@etagerca.ru'>sale@etagerca.ru</a>
                        </div>
                    </div>
                </div>

                <div className='other_info'>
                    <p>
                        Все ресурсы сайта ???, включая текстовую, графическую и видео информацию, структуру и оформление страниц, защищены российскими и международными законами и соглашениями об охране авторских прав и интеллектуальной собственности (статьи 1259 и 1260 главы 70 "Авторское право" Гражданского Кодекса Российской Федерации от 18 декабря 2006 года N 230-ФЗ).
                    </p>

                    <div className='copyright' itemScope itemType='http://schema.org/WPFooter'>
                        <p>
                            &copy;
                            <span itemProp='copyrightYear'> 2012–2024</span>
                            <span itemProp='copyrightHolder'> Интернет–магазин дизайнерской мебели и декора</span>
                        </p>
                    </div>

                    {/* <div className='extra'>
                        <a href='/policy/' target='_blank'>Политика конфиденциальности</a>
                        <a href='/sitemap/' target='_blank'>Карта сайта</a>
                    </div> */}
                </div>
            </footer>
        );
    }
}