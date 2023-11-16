import './App.css';
import React, { useEffect, useState } from 'react';

function App() {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        // Make an API request using the fetch function
        const response = await fetch('http://localhost:8080/users');
        
        // Check if the request was successful (status code 200)
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }

        // Parse the JSON data from the response
        const result = await response.json();

        // Update state with the fetched data
        setData(result);
      } catch (error) {
        // Handle any errors that occurred during the fetch
        setError(error);
      } finally {
        // Set loading to false after the request is complete, regardless of success or failure
        setLoading(false);
      }
    };

    // Call the fetchData function when the component mounts
    fetchData();
  }, []); 

  console.log({data})

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }
  return (
    <div className="App">
      <h1>Here is the user list LOL</h1>
      {data?.map((item)=>{
        return (
          <>
          <p>{item.id}</p>
          <p>{item.name}</p>
          <p>{item.email}</p>
          </>
        )
      })}
    </div>
  );
}

export default App;
