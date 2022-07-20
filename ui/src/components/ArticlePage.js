import './Articles.css';
import {useEffect, useState} from "react";
import consts from "../consts";
import Button from "./Button";

function ArticlePage({articleMeta}) {

    const [article, setArticle] = useState({})
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState(null);

    useEffect(() => {
        fetch(`${consts.API_ADDRESS}/articles/${articleMeta.ID}`)
            .then(res => res.json())
            .then(
                (result) => {
                    setArticle(result.response);
                    setIsLoading(false);
                }, (error) => {
                    setError(error)
                    setIsLoading(false);
                }
            )
    }, [])

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (isLoading) {
        return <div>Loading...</div>;
    } else {
        return (
            <div className="article_page">
                <h2 className="article__title">{articleMeta["title"]}</h2>
                <div className="article__section">
                    <span>By {articleMeta["author"]}</span>
                    <span>Published at {new Date(articleMeta["published_at"]).toDateString()}</span>
                    <span>Source: {articleMeta["source"]}</span>
                    <span>Pages: {articleMeta["pages_count"]}</span>
                </div>
                <section className="article__buttons">
                    <Button>Content</Button>
                    <Button>Index</Button>
                    <Button>Linguistic Expression</Button>
                </section>
                <p className="article__content">{article["content"]}</p>
            </div>
        )
    }
}

export default ArticlePage;