import '../../styles/Articles.css';
import consts from "../../consts";
import useFetch from "../../custom_hooks/useFetch";

function ArticleLing({articleId}) {
    const [article, isLoading, error] = useFetch(`${consts.API_ADDRESS}/articles/${articleId}`)

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (isLoading) {
        return <div>Loading...</div>;
    } else {
        return (
            <div>Ling</div>
        )
    }
}

export default ArticleLing;