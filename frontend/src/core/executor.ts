interface CommandDefinition {
  description: string
  usage?: string
  handler: (terminal: promptly.TerminalType, args: string[]) => void | Promise<void>
}

interface ParsedCommand {
  command: string
  args: string[]
  rawArgs: string
}

class CommandParser {
  private commands = new Map<string, CommandDefinition>()

  /**
   * 注册命令
   */
  register(name: string, definition: CommandDefinition): void {
    this.commands.set(name.toLowerCase(), definition)
  }

  /**
   * 解析命令行输入
   */
  parse(input: string): ParsedCommand | null {
    const trimmed = input.trim()
    if (!trimmed) return null

    const tokens = this.tokenize(trimmed)
    if (tokens.length === 0) return null

    return {
      command: tokens[0]?.toLowerCase() || '',
      args: tokens.slice(1),
      rawArgs: this.extractRawArgs(trimmed)
    }
  }

  /**
   * 执行命令
   */
  async execute(input: string, terminal: promptly.TerminalType): Promise<void> {
    const parsed = this.parse(input)
    if (!parsed) return

    const command = this.commands.get(parsed.command)
    if (!command) {
      terminal.error(`promptly: command not found - [${parsed.command}]`)
      return
    }

    try {
      await command.handler(terminal, parsed.args)
    } catch (error) {
      terminal.error(`Error executing command: ${error}`)
    }
  }

  /**
   * 获取所有已注册的命令
   */
  getCommands(): Array<{ name: string; definition: CommandDefinition }> {
    return Array.from(this.commands.entries()).map(([name, definition]) => ({
      name,
      definition
    }))
  }

  /**
   * 词法分析 - 支持引号和转义字符
   */
  private tokenize(input: string): string[] {
    const tokens: string[] = []
    let current = ''
    let inQuotes = false
    let quoteChar = ''
    let escapeNext = false
    let i = 0

    while (i < input.length) {
      const char = input[i]

      if (escapeNext) {
        current += char
        escapeNext = false
        i++
        continue
      }

      if (char === '\\') {
        escapeNext = true
        i++
        continue
      }

      if (!inQuotes && (char === '"' || char === "'")) {
        inQuotes = true
        quoteChar = char
        i++
        continue
      }

      if (inQuotes && char === quoteChar) {
        inQuotes = false
        quoteChar = ''
        i++
        continue
      }

      if (char != undefined && !inQuotes && /\s/.test(char)) {
        if (current) {
          tokens.push(current)
          current = ''
        }
        i++
        continue
      }

      current += char
      i++
    }

    if (current) {
      tokens.push(current)
    }

    if (inQuotes) {
      throw new Error('Unclosed quote')
    }

    return tokens
  }

  /**
   * 提取原始参数字符串
   */
  private extractRawArgs(input: string): string {
    const firstSpace = input.indexOf(' ')
    return firstSpace === -1 ? '' : input.slice(firstSpace + 1)
  }
}

// 创建命令解析器实例
const parser = new CommandParser()

// 注册内置命令
parser.register('help', {
  description: 'Print help message',
  usage: 'help [command]',
  handler: (terminal, args) => {
    if (args.length > 0) {
      const commandName = args[0]
      if (commandName) {
        commandName.toLowerCase()
      }
      const command = parser.getCommands().find((cmd) => cmd.name === commandName)

      if (command) {
        let help = `Command: ${command.name}\n`
        help += `Description: ${command.definition.description}\n`
        if (command.definition.usage) {
          help += `Usage: ${command.definition.usage}`
        }
        terminal.info(help)
      } else {
        terminal.error(`No help available for command: ${commandName}`)
      }
    } else {
      const commands = parser.getCommands()
      let msg = 'Promptly is a terminal like web index.\n'
      msg += 'There are commands that supported:\n'

      commands.forEach(({ name, definition }) => {
        msg += `${name.padEnd(12)}---- ${definition.description}\n`
      })

      terminal.info(msg)
    }
  }
})

parser.register('clear', {
  description: 'Clear the terminal screen',
  handler: (terminal) => {
    terminal.clear()
  }
})

parser.register('echo', {
  description: 'Echo arguments to the terminal',
  usage: 'echo [args...]',
  handler: (terminal, args) => {
    if (args.length === 0) {
      terminal.info('')
      return
    }
    terminal.info(args.join(' '))
  }
})

parser.register('date', {
  description: 'Display current date and time',
  handler: (terminal) => {
    const now = new Date()
    terminal.info(now.toLocaleString())
  }
})

parser.register('history', {
  description: 'Display command history',
  handler: (terminal, args) => {
    const history = CommandHistory.get()
    if (history.length === 0) {
      terminal.info('No command history')
      return
    }
    let arg = args[0]
    if (!arg) {
      arg = '10'
    }
    const start = args.length > 0 ? parseInt(arg) || 0 : 0
    const end = Math.min(start + 10, history.length)

    history.slice(start, end).forEach((cmd, index) => {
      terminal.info(`${start + index + 1}: ${cmd}`)
    })
  }
})

class CommandHistory {
  private static history: string[] = []
  private static maxHistory = 100

  static add(command: string): void {
    if (command.trim()) {
      this.history.push(command)
      if (this.history.length > this.maxHistory) {
        this.history.shift()
      }
    }
  }

  static get(): string[] {
    return [...this.history]
  }

  static clear(): void {
    this.history = []
  }
}

// 导出执行函数
export const executeCommand = async (input: string, terminal: promptly.TerminalType) => {
  // 添加到历史记录
  CommandHistory.add(input)

  // 执行命令
  await parser.execute(input, terminal)
}

// 导出命令解析器和历史记录类，供其他模块使用
export { CommandParser, CommandHistory }
