import Navbar from "../components/navbar"
import PieChart from "../components/piechart"

function Dashboard() {
  return (
    <>
    <Navbar />
    <div class="absolute inset-y-0 left-72 w-full">
      <main class="p-1 h-auto w-9/12">
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
          <div class="flex items-center justify-center flex-col border-2 border-dashed border-gray-300 rounded-lg dark:border-gray-600 h-32 md:h-64">
            <h1 class="text-xl font-bold text-center py-2 bg-slate-400 rounded-t-lg w-full">
              Quantidade Total
            </h1>
            <p class="text-center hover:border-blue-200 border-2 w-full h-full pt-16 text-4xl font-extrabold">
              {Math.floor(Math.random() * 100)}
            </p>
          </div>
          <div class="flex items-center justify-center flex-col border-2 border-dashed border-gray-300 rounded-lg dark:border-gray-600 h-32 md:h-64">
            <h1 class="text-xl font-bold text-center py-2 bg-slate-400 rounded-t-lg w-full">
              Recebido Ultimo Mês
            </h1>
            <p class="text-center hover:border-blue-200 border-2 w-full h-full pt-16 text-4xl font-extrabold">
              R$ {(Math.random() *10).toFixed(2)}
            </p>
          </div>
          <div class="flex items-center justify-center flex-col border-2 border-dashed border-gray-300 rounded-lg dark:border-gray-600 h-32 md:h-64">
            <h1 class="text-xl font-bold text-center py-2 bg-slate-400 rounded-t-lg w-full">
              Variação
            </h1>
            <p class="text-center text-green-500  hover:border-blue-200 border-2 w-full h-full pt-16 text-4xl font-extrabold">
              + {(Math.random() *10).toFixed(2)} %
            </p>
          </div>
          <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-32 md:h-64" />
        </div>
        <div class="border-2 border-dashed rounded-lg bg-blue-200 border-gray-300 dark:border-gray-600 h-96 mb-4">
          <PieChart />
        </div>
        <div class="grid grid-cols-2 gap-4 mb-4">
          <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-48 md:h-72" />
          <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-48 md:h-72" />
          <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-48 md:h-72" />
          <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-48 md:h-72" />
        </div>
        <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-96 mb-4" />
        <div class="grid grid-cols-2 gap-4">
          <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-48 md:h-72" />
          <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-48 md:h-72" />
          <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-48 md:h-72" />
          <div class="border-2 border-dashed rounded-lg border-gray-300 dark:border-gray-600 h-48 md:h-72" />
        </div>
      </main>
    </div>
  </>
  )

}

export default Dashboard