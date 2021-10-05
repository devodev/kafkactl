package config

import (
	"fmt"
)

type Context struct {
	Name      string `yaml:"name"`
	BaseURL   string `yaml:"baseURL"`
	ClusterID string `yaml:"clusterID"`
}

type ConfigFile struct {
	CurrentContext string     `yaml:"currentContext"`
	Contexts       []*Context `yaml:"contexts"`
}

type Config struct {
	currentContext string
	contexts       map[string]*Context
}

func New() *Config {
	return &Config{contexts: make(map[string]*Context)}
}

func newFromFile(f *ConfigFile) (*Config, error) {
	cfg := New()
	cfg.currentContext = f.CurrentContext

	for idx, ctx := range f.Contexts {
		if ctx == nil {
			continue
		}
		if ctx.Name == "" {
			return nil, fmt.Errorf("contexts[%d]: name not set", idx)
		}
		if _, ok := cfg.contexts[ctx.Name]; ok {
			return nil, fmt.Errorf("contexts[%d]: name '%s' already exists", idx, ctx.Name)
		}
		cfg.contexts[ctx.Name] = ctx
	}

	return cfg, nil
}

func (c Config) GetCurrentContextName() string {
	return c.currentContext
}

func (c Config) GetCurrentContext() (*Context, error) {
	if c.currentContext == "" {
		return nil, fmt.Errorf("currentContext is not set")
	}

	ctx, ok := c.contexts[c.currentContext]
	if !ok {
		return nil, fmt.Errorf("currentContext %s not found", c.currentContext)
	}
	return ctx, nil
}

func (c *Config) SetCurrentContext(name string) error {
	if _, err := c.GetContext(name); err != nil {
		return err
	}
	c.currentContext = name
	return nil
}

func (c Config) GetContext(name string) (*Context, error) {
	ctx, ok := c.contexts[name]
	if !ok {
		return nil, fmt.Errorf("context with name '%s' not found", name)
	}
	return ctx, nil
}

func (c Config) GetContexts() []*Context {
	contexts := make([]*Context, 0, len(c.contexts))
	for _, ctx := range c.contexts {
		contexts = append(contexts, ctx)
	}
	return contexts
}

func (c *Config) AddContext(ctx *Context) error {
	if _, err := c.GetContext(ctx.Name); err == nil {
		return fmt.Errorf("context with name '%s' already exists", ctx.Name)
	}
	c.contexts[ctx.Name] = ctx
	return nil
}

func (c *Config) UpdateContext(ctx *Context) error {
	if _, err := c.GetContext(ctx.Name); err != nil {
		return err
	}
	c.contexts[ctx.Name] = ctx
	return nil
}

func (c *Config) RemoveContext(name string) error {
	if _, err := c.GetContext(name); err != nil {
		return err
	}
	delete(c.contexts, name)
	return nil
}

func (c *Config) configFile() *ConfigFile {
	contexts := c.GetContexts()
	cfgFile := &ConfigFile{
		CurrentContext: c.currentContext,
		Contexts:       contexts,
	}
	return cfgFile
}
