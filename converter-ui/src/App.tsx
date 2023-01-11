import * as core from '@mantine/core'
import './App.css'
import Converter from './Converter'

function App() {
  return (
    <core.MantineProvider withGlobalStyles withNormalizeCSS>
      <Converter />
    </core.MantineProvider>
  )
}

export default App
