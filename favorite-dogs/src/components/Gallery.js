import React, { useState, useEffect } from 'react';

const Gallery = () => {
  const [dogs, setDogs] = useState([]);
  const [favorites, setFavorites] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    console.log('useEffect: gallery')
    const storedFavorites = localStorage.getItem('favorites');
    if (storedFavorites) {
      setFavorites(JSON.parse(storedFavorites));
    }
    fetchDogs();
  }, []);

  const fetchDog = async () => {
    try {
      const response = await fetch('https://random.dog/woof.json');
      const data = await response.json();
      if(data.url.match(/\.(jpeg|jpg|gif|png)$/i)) {
        return data.url;
      } else {
        return fetchDog();
      }
    } catch (error) {
      console.error('Error fetching dog:', error);
    }
  }
  const fetchDogs = async () => {
    setLoading(true);
    try {
      const imageUrls = await Promise.all([
        fetchDog(),
        fetchDog(),
        fetchDog(),
        fetchDog(),
        fetchDog(),
        fetchDog(),
      ]);
  
      setDogs(imageUrls);
    } catch (error) {
      console.error('Error fetching dogs:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleFavorite = (dog) => {
    setFavorites((prevFavorites) => [...prevFavorites, dog]);
    localStorage.setItem('favorites', JSON.stringify([...favorites, dog]));
  };

  const handleUnfavorite = (dog) => {
    setFavorites((favorites) => {
      const updatedFavorites = favorites.filter((favorite) => favorite !== dog);
      localStorage.setItem('favorites', JSON.stringify(updatedFavorites));
      return updatedFavorites;
    });
  };

  const handleRefresh = () => {
    fetchDogs();
  };

  return (
    <div className="container">
      <div className="row">
        <div className="col-12 d-flex justify-content-center">
          <h2 className="p-3">Gallery</h2>
        </div>
      </div>
      <div className="row">
        <div className="col-12 d-flex justify-content-center">
          <button className="btn btn-primary mb-4" onClick={handleRefresh}>
            Refresh
          </button>
        </div>
      </div>
      <div className="row">
        {loading ? (
          <p>Loading...</p>
        ) : (
          dogs.map((dog, index) => (
            <div key={index} className="col-md-4 col-sm-6 mb-4">
              <div className="card h-100">
                <img src={dog} alt="Dog" className="card-img-top" />
                <div className="card-body d-flex flex-column">
                <button
                    className={`btn ${favorites.includes(dog) ? 'btn-danger' : 'btn-primary'} mt-auto`}
                    onClick={() => favorites.includes(dog) ? handleUnfavorite(dog) : handleFavorite(dog)}
                  >
                    {favorites.includes(dog) ? 'Unfavorite' : 'Favorite'}
                  </button>
                </div>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default Gallery;
