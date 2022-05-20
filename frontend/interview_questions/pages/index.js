import { useCallback, useState } from 'react'

export default function Home() {

  const [showAnswer, setShowAnswer] = useState(false)

  const [randomQuestion, setRandomQuestion] = useState(
    {
      "Question": "",
      "Answer": "",
      "Category": "",
      "Level": "",
      "Source": ""
    }
  )

  const toggleAnswer = useCallback(() => {
    setShowAnswer(!showAnswer)
  },[showAnswer])

  const getNewQuestion = useCallback(async () => {
    fetch('http://localhost:8080/questions')
      .then((rsp) => rsp.json())
      .then((data) => {
        setShowAnswer(false)
        setRandomQuestion(data[Math.floor(Math.random() * (data.length-1))])
      })
  },[])

  return (
  <div>
    <p id='question'>
      {randomQuestion['Question']}
    </p>
    {showAnswer &&
    <p id='answer'>
      {randomQuestion['Answer']}
    </p>
    }
    <button onClick={getNewQuestion}>Get new question</button>
    <button onClick={toggleAnswer}>{`${showAnswer ? "Hide" : "Show"} Answer`}</button>
  </div>
  )
}
