export function formatDate(date: Date): string {
  const day = date.getDate();
  const month = date.toLocaleString('default', { month: 'long' });
  const year = date.getFullYear();

  function getDaySuffix(day: number): string {
    if (day > 3 && day < 21) return 'th'; // 4th to 20th
    switch (day % 10) {
      case 1:
        return 'st';
      case 2:
        return 'nd';
      case 3:
        return 'rd';
      default:
        return 'th';
    }
  }

  const daySuffix = getDaySuffix(day);

  return `${day}<sup>${daySuffix}</sup> ${month} ${year}`;
}

export const multiplyValues = (input: string) => {
  const [value1, value2] = input.split('x').map(Number);
  return value1 * value2;
};

export const formatNumbers = (value: number): number => {
  return Math.trunc(value * 100) / 100;
};
