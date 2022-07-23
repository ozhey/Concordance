import '../../styles/Index.css';
import Button from "../Button";
import {useState} from "react";

function IndexWord({wordObj, selectWord}) {
    const [expand, setExpand] = useState(false)
    let wordsIndex = []
    if (expand) {
        const occurrences = wordObj["index"].split('\n')
        const indexTable = occurrences.map((occurrence) => {
            const pos = occurrence.split(',')
            const [articleId, page, line, word] = pos
            return <div onClick={() => selectWord({pos: {articleId, page, line, word}, word: wordObj["word"]})} key={occurrence} className="index__word__row index__word__row--clickable">
                <span>{articleId}</span>
                <span>{page}</span>
                <span>{line}</span>
                <span>{word}</span>
            </ div>
        })
        wordsIndex =
            <div className="index__word__table">
                <div className="index__word__row">
                    <span>Article ID</span>
                    <span>Page</span>
                    <span>Line</span>
                    <span>Word</span>
                </div>
                    {indexTable}
            </div>
    }

    function expandOrCollapse() {
        setExpand((prev) => !prev)
    }


    return (
        <div className="index__word">
            <b>Word</b>
            <b>Occurrences</b>
            <div style={{gridColumnStart: "3", gridRow: "1 / 3"}}>
                <Button onClick={expandOrCollapse} size="small">{expand ? `Hide Index` : `Show Index`}</Button>
            </div>
            <div>{wordObj["word"]} </div>
            <div>{wordObj["count"]}</div>
            {expand ?
                <div style={{gridColumn: "1 / span 3", width: "100%"}}>
                    {wordsIndex}
                </div> :
                null
            }
        </div>
    )
}

export default IndexWord;