package database

const (
	getArticleByID = `
SELECT string_agg(article_lines.line, E'\n') as article,
       AVG(words_in_line)                    AS words_in_line,
       AVG(chars_in_word)                    AS chars_in_word,
       AVG(pages_num)                        AS pages_num
FROM (SELECT string_agg(word, ' ' ORDER BY word_number) AS line,
             AVG(a.pages_count)                         AS pages_num,
             AVG(al.word_count)                         AS words_in_line,
             AVG(char_count)                            AS chars_in_word
      FROM article_words
               JOIN article_lines al ON al.id = article_words.article_line_id
               JOIN article_pages ap ON al.article_page_id = ap.id
               JOIN articles a ON ap.article_id = a.id
      WHERE a.id = ?
      GROUP BY page_number, line_number
      ORDER BY page_number, line_number) AS article_lines`

	getArticles = `
SELECT *
FROM articles
`

	getArticleWords = `
SELECT word, COUNT(word),
       string_agg(CONCAT('page: ',ap.page_number::text, ' line: ',al.line_number::text, ' word: ',word_number::text),  E'\n' ORDER BY ap.page_number, al.line_number, word_number) AS index
FROM article_words
         JOIN article_lines al on al.id = article_words.article_line_id
         JOIN article_pages ap on ap.id = al.article_page_id
         JOIN articles a on a.id = ap.article_id
WHERE a.id = ?
GROUP BY word
ORDER BY COUNT(word) DESC`
)
