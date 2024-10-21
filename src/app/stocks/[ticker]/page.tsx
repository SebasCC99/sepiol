export default async function Page({params}: {params : {ticker: string}}) {
  fetch(`http://localhost:8080/stocks/${params.ticker}`);

  return (
    <main>      
      <h1>Ticker: {params.ticker.toUpperCase()}</h1>
    </main>
  );
}
