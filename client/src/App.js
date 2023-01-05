import React,{useEffect, useState} from 'react'
import {Card} from 'antd'
function App() {
  const [data,setData] = useState([])
  useEffect(()=>{
    fetchData()
  },[])
  const fetchData = async()=>{
    const responseData = await fetch('http://127.0.0.1:4000/api/todos/list')
    console.log(responseData)
    setData(responseData)
  }
  return (
    <>
      Hi
    </>
  );
}

export default App;
