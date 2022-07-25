import '../../styles/Index.css';
import consts from "../../consts";
import useFetch from "../../api/useFetch";
import IndexWord from "./IndexWord";
import {useState} from "react";
import ContextWindow from "../ContextWindow";

function ArticleIndex({articleId}) {
    const [index, isLoading, error] = useFetch(`${consts.API_ADDRESS}/article_words/index?article_id=${articleId}&word_group_id=${""}`, null, [])
    const [selectedWord, setSelectedWord] = useState(null)

    const words = index.map((wordObj) => <IndexWord key={wordObj["word"]} wordObj={wordObj} selectWord={setSelectedWord}/>)

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (isLoading) {
        return <div>Loading...</div>;
    } else {
        return (
            <section>
                {(selectedWord !== null) ?
                    <ContextWindow pos={selectedWord["pos"]} expr={selectedWord["word"]} /> :
                    ""
                }
                <div className="index">
                    {words}
                </div>
            </section>
        )
    }
}

export default ArticleIndex;