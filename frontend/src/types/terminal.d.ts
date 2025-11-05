declare namespace promptly {
  type OutputStatusType = 'info' | 'success' | 'warning' | 'error' | 'system'

  interface TerminalType {
    print(output: Omit<promptly.OutputType, string>)
    printCommand(command: string, result: promptly.OutputType[] = [])
    clear()
    error(message: string)
    success(message: string)
    warning(message: string)
    info(message: string)
  }

  interface OutputType {
    type: 'command' | 'text' | 'component'
    text?: string
    result_list?: OutputType[]
    component?: any
    status: OutputStatusType
    props?: any
    collapsible?: boolean
  }

  interface UserType {
    prompt?: string
    username?: string
  }
}
