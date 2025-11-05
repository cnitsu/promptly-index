declare namespace promptly {
  type OutputStatusType = 'info' | 'success' | 'warning' | 'error' | 'system'
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
