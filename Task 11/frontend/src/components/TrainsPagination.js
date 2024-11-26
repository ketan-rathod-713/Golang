import React, { useEffect, useState } from 'react'
import Pagination from './Pagination';
import sortup from "../icons/sortup.svg"
import sortdown from "../icons/sortdown.svg"

const TrainsPagination = () => {
    const [page, setPage] = useState(1)
    const [limit, setLimit] = useState(10)
    const [total, setTotal] = useState(10)
    const [sort, setSort] = useState("number")
    const [order, setOrder] = useState(1)
    const [searchText, setSearchText] = useState("")
    const [trains, setTrains] = useState([])
    const [loadData, setLoadData] = useState(true)

    const toggleloaddata = ()=>{
        if(loadData){
            setLoadData(false)
        } else {
            setLoadData(true)
        }
    }

    const handlePageChange = (newPage)=>{
        setPage(newPage)
    }

    const resetSettings = (newPage)=>{
        // setTrains([])
        setPage(1)
        setLimit(10)
        setOrder(1)
        setSearchText("")
        toggleloaddata()
    }

    const handleFindTrainInputChange = (event)=>{
        setSearchText(event.target.value)
        console.log(event.target.value)
    }

    const handleLimitChange = (newLimit)=>{
        setPage(1)
        setLimit(newLimit)
    }

    const sortHandler = (sortField) =>{
        console.log(sortField, sort)
        if(sort == sortField){
            // toggle order 
            if(order == 1) {
                setOrder(-1)
            } else {
                setOrder(1)
            }
        } else {
            setSort(sortField)
            setOrder(1)
        }
    } 

    useEffect(()=>{
        // before requesting 
        console.log("state before request", page, limit, sort, order, searchText)
        fetch(`http://localhost:8080/train?page=${page}&limit=${limit}&sort=${sort}&order=${order}&search=${searchText}`).then(response => {
            return response.json()
        }).then(response => {
            if(response.status){
                alert(response.message)
                return
            } else {
                console.log(response)
                setTrains(response.data)
                setTotal(response.total)
            }
            
        }).catch(response => {
            alert("failed to fetch train records")
        })
    },[page, limit, sort, order, loadData])

  return (
    <div className='table-wrapper'>
        <div className='find-trains-wrapper'>
            <h1>Find Trains</h1>
            <div className='find-trains-input'>
                <input type='text' onChange={handleFindTrainInputChange} value={searchText} placeholder='Search By Train Station Name'/>
            </div>
            <button onClick={toggleloaddata} className='find-train-btn'>FIND</button>
        </div>

        <div className='main-header'>
        <div className='records-per-page'>
            <label>Records Per Page</label>
            <div>
                <select id='select' value={limit} onChange={(e)=>handleLimitChange(e.target.value)}>
                <option value={10}>10</option>
                <option value={50}>50</option>
                <option value={100}>100</option>
            </select>
            </div>
        </div>
        <button onClick={resetSettings}>RESET SETTINGS</button>
        </div>


        <div className='table-header'>
            <h4> Page {page} out of {Math.ceil(total/limit)} Pages</h4>
            <h4> Total Trains {total}</h4>
        </div>
        <table>
            <thead>
                <tr>
                    <th onClick={(e)=> sortHandler("number")}>Number</th>
                    <th onClick={(e)=> sortHandler("name")}>Name</th>
                    <th onClick={(e)=> sortHandler("source")}>Source</th>
                    <th onClick={(e)=> sortHandler("destination")}>Destination</th>
                </tr>
            </thead>
            <tbody>
                {trains.map(train => (
                    <tr key={train.number}>
                        <td>{train.number}</td>
                        <td>{train.name}</td>
                        <td>{train.source}</td>
                        <td>{train.destination}</td>
                    </tr>
                ))}
            </tbody>
        </table>
        <Pagination page={page} limit={limit} total={total} handlePageChange={handlePageChange}/>
    </div>
  )
}

export default TrainsPagination