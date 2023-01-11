import { Button, Center, Grid, NumberInput, Select, SelectItem, TextInput } from "@mantine/core";
import { useCallback, useEffect, useState } from "react";
import { API_URL } from "./consts";

interface State {
    res: string | undefined,
    currencies: SelectItem[],
    val: string | undefined,
    from: string | null,
    to: string | null
  }
  

const Converter = () => {    
    const [state, setState] = useState({ 
        currencies: [] as unknown,
    } as State)

    useEffect(() => {
        console.log("Hallo")
        fetch(`${API_URL}/currencies`)
            .then(res => res.json())
            .then(res => {
            setState({ ...state, currencies: res })
            })
    }, [])

    const onConvertClick = useCallback(() => {
        fetch(`${API_URL}/convert?from=${state.from?.toLowerCase()}&to=${state.to?.toLowerCase()}&val=${state.val}`)
            .then(res => res.json())
            .then(res => {
                setState({ ...state, res: res.value })
            })
    }, [state]);

    const setVal = useCallback((value: string) => {
        let tmp = value.replace(/[^\d\.]/g, '');
        let tmptmp = tmp.length ? tmp : "";

        setState({ ...state, val: tmptmp })
    }, [state]);

    return (
        <Center style={{ width: '100vw', height: '100vh' }}>
            <Grid>
                <Grid.Col span={2}></Grid.Col>
                <Grid.Col span={2}><Select placeholder="Please Select..." data={ state.currencies } onChange={ (e) => setState({ ...state, from: e }) } searchable /></Grid.Col>
                <Grid.Col span={2}><TextInput placeholder="123.2" value={ state.val } onChange={ (n) => setVal(n.currentTarget.value) } /></Grid.Col>
                <Grid.Col span={2}><TextInput value={ state.res } disabled /></Grid.Col>
                <Grid.Col span={2}><Select placeholder="Please Select..." data={ state.currencies } defaultValue="loading" onChange={ (e) => setState({ ...state, to: e }) } searchable /></Grid.Col>
                <Grid.Col span={2}></Grid.Col>
                <Grid.Col span={12}><Button mx="auto" display="block" onClick={ onConvertClick }>Convert</Button></Grid.Col>
            </Grid>
      </Center>
    )
}

export default Converter;