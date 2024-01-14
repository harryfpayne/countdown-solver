import React from 'react';

function App() {
  return (
    <div className="App bg-blue-300">
      <header className="h-60">
        Explainer
      </header>

      <section className="p-20 space-y-6">
        <h1>Numbers round</h1>
        <div className={"flex justify-center"}>
          <input
            className={"w-[4ch] text-5xl bg-blue-700 text-white text-center border-blue-500 border-4 font-bold"}
            type="text"
          />
        </div>
        <div className={"flex flex-row justify-center space-x-4"}>
          {[1, 2, 3, 4, 5, 6].map((n) => (
            <input
              key={`number-${n}`}
              className={"w-[3.5ch] h-[3ch] text-4xl bg-blue-700 text-white text-center border-blue-500 border-2 font-bold"}
              type="text"
            />
          ))}
        </div>
        <div className={"flex justify-center"}>
          <div className={"bg-gray-50 w-[600px] h-[400px] background-grid flex justify-center align-text-top p-4 font-bold leading-8"}>
            100 + 75 = 175 <br />
            175 * 6 = 1050 <br />
            1050 + 25 = 1075 <br />
            1050 / 5 = 210 <br /><br /><br />
            Found 1020 solutions in 28 seconds
          </div>
        </div>
      </section>

      <section className="p-20 space-y-6">
        <h1>Letters round</h1>
        <div className={"flex flex-row justify-center space-x-4"}>
          {[1, 2, 3, 4, 5, 6, 7, 8, 9].map((n) => (
            <input
              key={`letter-${n}`}
              className={"w-[3ch] h-[3ch] text-4xl bg-blue-700 text-white text-center border-blue-500 border-2 font-bold"}
              type="text"
            />
          ))}
        </div>
      </section>
    </div>
  );
}

export default App;
