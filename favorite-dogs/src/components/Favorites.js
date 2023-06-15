import React, { useEffect, useState } from 'react';

const Favorites = () => {
  const [favorites, setFavorites] = useState([]);

  useEffect(() => {
    console.log('useEffect: favouries dogs')
    const storedFavorites = localStorage.getItem('favorites');
    if (storedFavorites) {
      setFavorites(JSON.parse(storedFavorites));
    }
  }, []);

  const handleUnfavorite = (dog) => {
    setFavorites((favorites) => {
      const updatedFavorites = favorites.filter((favorite) => favorite !== dog);
      localStorage.setItem('favorites', JSON.stringify(updatedFavorites));
      return updatedFavorites;
    });
  };
  

  return (
    <div className="container">
      <div className="row">
        <div className="col-12 d-flex justify-content-center">
          <h2 className="p-3">Favorites</h2>
        </div>
      </div>
      <div className="row">
      {favorites.map((dog, index) => (
        <div key={index} className="col-md-4 col-sm-6 mb-4">
        <div className="card h-100">
          <img src={dog} alt="Dog" className="card-img-top" />
          <div className="card-body d-flex flex-column">
            <button className="btn btn-danger mt-auto" onClick={() => handleUnfavorite(dog)}>
              Unfavorite
            </button>
          </div>
        </div>
      </div>
      ))}
      </div>
    </div>
  );
};

export default Favorites;
