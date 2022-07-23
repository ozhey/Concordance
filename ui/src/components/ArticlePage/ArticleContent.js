import '../../styles/Articles.css';
import consts from "../../consts";
import useFetch from "../../custom_hooks/useFetch";

function ArticleContent({articleId}) {
    const [article, isLoading, error] = useFetch(`${consts.API_ADDRESS}/articles/${articleId}`)

    let lineElems
    if (article["content"]) {
        const articleLines = article["content"].split("\n")
        lineElems = articleLines.map((line, i) => <p key={i}><b style={{marginRight:"2px"}}>{i+1} </b>{line}</p>)
    }

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (isLoading) {
        return <div>Loading...</div>;
    } else {
        return <div className="article__content">
            {lineElems}
        </div>
    }
}

export default ArticleContent;