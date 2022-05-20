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

  const openQuestionSource = useCallback(() => {
    window.open(randomQuestion['Source'], "_blank")
  },[])

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
  <div className="flex justify-center">
    <div className="block p-6 rounded-lg shadow-lg bg-white max-w-sm">
      <h5 className="text-gray-900 text-xl leading-tight font-medium mb-2">{randomQuestion['Question']}</h5>
      <p className="text-gray-700 text-base mb-4">
        {showAnswer && <p className="text-grey-700 text-base">{randomQuestion['Answer']}</p>}
      </p>
      <button type="button" className=" inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out" onClick={getNewQuestion}>Get new question</button>
      <button type="button" className=" inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out" onClick={toggleAnswer}>{`${showAnswer ? "Hide" : "Show"} Answer`}</button>
      <button type="button" className=" inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out" onClick={openQuestionSource}>{`Source`}</button>
    </div>
  </div>
  )
}
