import './Articles.css';
import ArticleCard from "./ArticleCard";
import {useEffect, useState} from "react";
import consts from '../consts.js'
import ArticlePage from "./ArticlePage";

function Articles() {
    const [articles, setArticles] = useState([])
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState(null);
    const [selectedArticleID, setSelectedArticleID] = useState(0)

    useEffect(() => {
        fetch(`${consts.API_ADDRESS}/articles`)
            .then(res => res.json())
            .then(
                (result) => {
                    setArticles(result.response);
                    setIsLoading(false);
                }, (error) => {
                    setError(error)
                    setIsLoading(false);
                }
            )
    }, [])


    function handleCardClick(articleID) {
        setSelectedArticleID(articleID)
    }
    const articleCards = articles.map(article => <ArticleCard key={article.ID} article={article} onClick={() => handleCardClick(article.ID)} />)

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (isLoading) {
        return <div>Loading...</div>;
    } else if (selectedArticleID !== 0) {
        return <ArticlePage id={selectedArticleID}> </ArticlePage>
    } else {
        return (
            <section className="articles" >
                {articles.length ? articleCards : <div>No articles found</div>}
            </section>
        );
    }
}

export default Articles;