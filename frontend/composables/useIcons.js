export const useIcons = () => {
  const isIconReady = ref(false);

  // Delay icon rendering to avoid SSR issues
  if (process.client) {
    setTimeout(() => {
      isIconReady.value = true;
    }, 100);
  }

  const SafeIcon = (props) => {
    return h(Icon, {
      ...props,
      vIf: isIconReady.value,
    });
  };

  return {
    isIconReady,
    SafeIcon,
  };
};
